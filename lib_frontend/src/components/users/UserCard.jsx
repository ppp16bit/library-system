import React from 'react';
import './UserCard.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faEnvelope, faEdit, faTrashAlt } from '@fortawesome/free-solid-svg-icons';

const UserCard = ({ user, onEdit, onDelete }) => {
    return (
        <div className="user-card">
            <div className="card-field">
                <label><FontAwesomeIcon icon={faUser} className="field-icon" /> Nome:</label>
                <input type="text" value={user.name || ''} readOnly />
            </div>
            <div className="card-field">
                <label><FontAwesomeIcon icon={faEnvelope} className="field-icon" /> Email:</label>
                <input type="text" value={user.email || ''} readOnly />
            </div>

            <div className="card-actions">
                <button onClick={onEdit} className="edit-button">
                    <FontAwesomeIcon icon={faEdit} className="icon" /> Editar
                </button>
                <button onClick={onDelete} className="delete-button">
                    <FontAwesomeIcon icon={faTrashAlt} className="icon" /> Excluir
                </button>
            </div>
        </div>
    );
};

export default UserCard;