import React, { useState, useEffect } from 'react';
import './UserForm.css';

const UserForm = ({ userToEdit, onSubmit, onCancel }) => {
    const [user, setUser] = useState({
        name: '',
        email: ''
    });
    const [errors, setErrors] = useState({});

    useEffect(() => {
        if (userToEdit) {
            setUser(userToEdit);
        } else {
            setUser({ name: '', email: '' });
        }
        setErrors({});
    }, [userToEdit]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setUser(prevUser => ({ ...prevUser, [name]: value }));
        setErrors(prevErrors => ({ ...prevErrors, [name]: '' }));
    };

    const validate = () => {
        const newErrors = {};
        if (!user.name.trim()) newErrors.name = 'Nome é obrigatório.';
        if (!user.email.trim()) {
            newErrors.email = 'Email é obrigatório.';
        } else if (!/\S+@\S+\.\S+/.test(user.email)) {
            newErrors.email = 'Email inválido.';
        }
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (validate()) {
            onSubmit(user);
        }
    };

    return (
        <div className="user-form-container">
            <h2>{userToEdit ? 'Editar Usuário' : 'Adicionar Novo Usuário'}</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="userName">Nome:</label>
                    <input
                        type="text"
                        id="userName"
                        name="name"
                        value={user.name}
                        onChange={handleChange}
                        placeholder="Digite o nome do usuário"
                        className={errors.name ? 'input-error' : ''}
                    />
                    {errors.name && <p className="error-text">{errors.name}</p>}
                </div>
                <div className="form-group">
                    <label htmlFor="userEmail">Email:</label>
                    <input
                        type="email"
                        id="userEmail"
                        name="email"
                        value={user.email}
                        onChange={handleChange}
                        placeholder="Digite o email do usuário"
                        className={errors.email ? 'input-error' : ''}
                    />
                    {errors.email && <p className="error-text">{errors.email}</p>}
                </div>
                <div className="form-actions">
                    <button type="submit" className="submit-button">
                        {userToEdit ? 'Salvar Alterações' : 'Adicionar Usuário'}
                    </button>
                    <button type="button" onClick={onCancel} className="cancel-button">
                        Voltar para a Lista
                    </button>
                </div>
            </form>
        </div>
    );
};

export default UserForm;
