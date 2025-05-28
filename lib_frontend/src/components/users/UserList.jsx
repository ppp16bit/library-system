import React from 'react';
import UserCard from './UserCard';
import './UserList.css';

const UserList = ({ users, onEditUser, onDeleteUser }) => {
    if (!users) {
        return (
            <div className="loading-container">
                <p>Carregando usuários...</p>
            </div>
        );
    }

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
};

export default UserList;
