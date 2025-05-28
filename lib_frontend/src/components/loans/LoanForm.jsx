// src/components/LoanForm.jsx
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { API_BASE_URL } from "../../constants";
import './LoanForm.css';

const LoanForm = ({ loanToEdit, onSubmit, onCancel }) => {
    const [loan, setLoan] = useState({
        userId: '',
        bookId: '',
        loanDate: new Date().toISOString().substring(0, 16),
        returnDate: '',
        returned: false
    });
    const [users, setUsers] = useState([]);
    const [books, setBooks] = useState([]);
    const [errors, setErrors] = useState({});
    const [loadingData, setLoadingData] = useState(true);

    useEffect(() => {
        const fetchUsersAndBooks = async () => {
            try {
                const [usersResponse, booksResponse] = await Promise.all([
                    axios.get(`${API_BASE_URL}/users`),
                    axios.get(`${API_BASE_URL}/books`)
                ]);
                setUsers(usersResponse.data);
                setBooks(booksResponse.data);
            } catch (err) {
                console.error("Erro ao carregar usuários ou livros:", err);
                setErrors({ general: "Não foi possível carregar usuários ou livros para o formulário." });
            } finally {
                setLoadingData(false);
            }
        };
        fetchUsersAndBooks();
    }, []);

    useEffect(() => {
        if (loanToEdit) {
            setLoan({
                ...loanToEdit,
                loanDate: loanToEdit.loanDate ? new Date(loanToEdit.loanDate).toISOString().substring(0, 16) : '',
                returnDate: loanToEdit.returnDate ? new Date(loanToEdit.returnDate).toISOString().substring(0, 16) : ''
            });
        } else {
            setLoan({
                userId: '',
                bookId: '',
                loanDate: new Date().toISOString().substring(0, 16),
                returnDate: '',
                returned: false
            });
        }
        setErrors({});
    }, [loanToEdit]);

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setLoan(prevLoan => ({
            ...prevLoan,
            [name]: type === 'checkbox' ? checked : value
        }));
        setErrors(prevErrors => ({ ...prevErrors, [name]: '' }));
    };

    const validate = () => {
        const newErrors = {};
        if (!loan.userId) newErrors.userId = 'Usuário é obrigatório.';
        if (!loan.bookId) newErrors.bookId = 'Livro é obrigatório.';
        if (!loan.loanDate) newErrors.loanDate = 'Data do empréstimo é obrigatória.';
        if (loan.loanDate && loan.returnDate && new Date(loan.returnDate) < new Date(loan.loanDate)) {
            newErrors.returnDate = 'Data de devolução não pode ser antes da data de empréstimo.';
        }
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (validate()) {
            const loanDataToSend = {
                ...loan,
                loanDate: loan.loanDate ? new Date(loan.loanDate).toISOString() : '',
                returnDate: loan.returnDate ? new Date(loan.returnDate).toISOString() : null
            };
            onSubmit(loanDataToSend);
        }
    };

    if (loadingData) {
        return (
            <div className="loan-form-container">
                <p style={{ textAlign: 'center', color: '#555' }}>Carregando opções de usuários e livros...</p>
            </div>
        );
    }

    return (
        <div className="loan-form-container">
            <h2>{loanToEdit ? 'Editar Empréstimo' : 'Registrar Novo Empréstimo'}</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="userId">Usuário:</label>
                    <select
                        id="userId"
                        name="userId"
                        value={loan.userId}
                        onChange={handleChange}
                        className={errors.userId ? 'input-error' : ''}
                    >
                        <option value="">Selecione um usuário</option>
                        {users.map(user => (
                            <option key={user.id} value={user.id}>{user.name}</option>
                        ))}
                    </select>
                    {errors.userId && <p className="error-text">{errors.userId}</p>}
                </div>
                <div className="form-group">
                    <label htmlFor="bookId">Livro:</label>
                    <select
                        id="bookId"
                        name="bookId"
                        value={loan.bookId}
                        onChange={handleChange}
                        className={errors.bookId ? 'input-error' : ''}
                    >
                        <option value="">Selecione um livro</option>
                        {books.map(book => (
                            <option key={book.id} value={book.id}>{book.title}</option>
                        ))}
                    </select>
                    {errors.bookId && <p className="error-text">{errors.bookId}</p>}
                </div>
                <div className="form-group">
                    <label htmlFor="loanDate">Data do Empréstimo:</label>
                    <input
                        type="datetime-local"
                        id="loanDate"
                        name="loanDate"
                        value={loan.loanDate}
                        onChange={handleChange}
                        className={errors.loanDate ? 'input-error' : ''}
                    />
                    {errors.loanDate && <p className="error-text">{errors.loanDate}</p>}
                </div>
                <div className="form-group">
                    <label htmlFor="returnDate">Data de Devolução (Opcional):</label>
                    <input
                        type="datetime-local"
                        id="returnDate"
                        name="returnDate"
                        value={loan.returnDate}
                        onChange={handleChange}
                        className={errors.returnDate ? 'input-error' : ''}
                    />
                    {errors.returnDate && <p className="error-text">{errors.returnDate}</p>}
                </div>
                <div className="form-group form-checkbox-group">
                    <input
                        type="checkbox"
                        id="returned"
                        name="returned"
                        checked={loan.returned}
                        onChange={handleChange}
                    />
                    <label htmlFor="returned">Devolvido?</label>
                </div>
                <div className="form-actions">
                    <button type="submit" className="submit-button">
                        {loanToEdit ? 'Salvar Alterações' : 'Registrar Empréstimo'}
                    </button>
                    <button type="button" onClick={onCancel} className="cancel-button">
                        Voltar para a Lista
                    </button>
                </div>
            </form>
        </div>
    );
};

export default LoanForm;
