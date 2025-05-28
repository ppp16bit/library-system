import React, { useEffect, useState, forwardRef, useImperativeHandle } from 'react';
import axios from 'axios';
import LoanCard from './LoanCard';
import { API_BASE_URL } from "../../constants";
import './LoanList.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';

const LoanList = forwardRef(({ onEditLoan, onDeleteLoan }, ref) => {
    const [loans, setLoans] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const fetchLoans = async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await axios.get(`${API_BASE_URL}/loans`);
            setLoans(response.data);
        } catch (err) {
            console.error("Erro ao buscar empréstimos:", err);
            setError("Não foi possível carregar os empréstimos. Tente novamente mais tarde.");
        } finally {
            setLoading(false);
        }
    };

    useImperativeHandle(ref, () => ({
        fetchLoans
    }));

    useEffect(() => {
        fetchLoans();
    }, []);

    if (loading) {
        return (
            <div className="loading-container">
                <FontAwesomeIcon icon={faSpinner} spin size="2x" color="#f0ad4e" />
                <p>Carregando empréstimos...</p>
            </div>
        );
    }

    if (error) return <p className="error-message">{error}</p>;
    if (loans.length === 0) return <p className="no-items-message">Nenhum empréstimo cadastrado.</p>;

    return (
        <div className="loan-list-container">
            {loans.map(loan => (
                <LoanCard
                    key={loan.id}
                    loan={loan}
                    onEdit={() => onEditLoan(loan)}
                    onDelete={() => onDeleteLoan(loan.id)}
                />
            ))}
        </div>
    );
});

export default LoanList;
