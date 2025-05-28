import React, { useEffect, useState, forwardRef, useImperativeHandle } from 'react';
import axios from 'axios';
import UserCard from './UserCard';
import { API_BASE_URL } from "../../constants";
import './UserList.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';

const UserList = forwardRef(({ onEditUser, onDeleteUser }, ref) => {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const fetchUsers = async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await axios.get(`${API_BASE_URL}/users`);
            setUsers(response.data);
        } catch (err) {
            console.error("Erro ao buscar usuários:", err);
            setError("Não foi possível carregar os usuários. Tente novamente mais tarde.");
        } finally {
            setLoading(false);
        }
    };

    useImperativeHandle(ref, () => ({
        fetchUsers
    }));

    useEffect(() => {
        fetchUsers();
    }, []);

    if (loading) {
        return (
            <div className="loading-container">
                <FontAwesomeIcon icon={faSpinner} spin size="2x" color="#6a5acd" />
                <p>Carregando usuários...</p>
            </div>
        );
    }

    if (error) return <p className="error-message">{error}</p>;
    if (users.length === 0) return <p className="no-items-message">Nenhum usuário cadastrado.</p>;

    return (
        <div className="user-list-container">
            {users.map(user => (
                <UserCard
                    key={user.id}
                    user={user}
                    onEdit={() => onEditUser(user)}
                    onDelete={() => onDeleteUser(user.id)}
                />
            ))}
        </div>
    );
});

export default UserList;
