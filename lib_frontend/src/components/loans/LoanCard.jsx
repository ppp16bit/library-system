import React from 'react';
import './LoanCard.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBook, faUser, faCalendarAlt, faCheckCircle, faTimesCircle, faEdit, faTrashAlt } from '@fortawesome/free-solid-svg-icons';

const LoanCard = ({ loan, onEdit, onDelete }) => {
    const formatDateTime = (isoString) => {
        if (!isoString) return 'N/A';
        const date = new Date(isoString);
        return new Intl.DateTimeFormat('pt-BR', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        }).format(date);
    };

    const statusColor = loan.returned ? '#28a745' : '#dc3545';

    return (
        <div className="loan-card">
            <div className="card-field">
                <label><FontAwesomeIcon icon={faUser} className="field-icon" /> Usuário:</label>
                <input type="text" value={loan.user ? loan.user.name : 'N/A'} readOnly />
            </div>
            <div className="card-field">
                <label><FontAwesomeIcon icon={faBook} className="field-icon" /> Livro:</label>
                <input type="text" value={loan.book ? loan.book.title : 'N/A'} readOnly />
            </div>
            <div className="card-field">
                <label><FontAwesomeIcon icon={faCalendarAlt} className="field-icon" /> Data do Empréstimo:</label>
                <input type="text" value={formatDateTime(loan.loanDate)} readOnly />
            </div>
            <div className="card-field">
                <label><FontAwesomeIcon icon={faCalendarAlt} className="field-icon" /> Data de Devolução:</label>
                <input type="text" value={formatDateTime(loan.returnDate)} readOnly />
            </div>
            <div className="card-field loan-status-field">
                <label>Status:</label>
                <span style={{ color: statusColor, fontWeight: 'bold' }}>
                    <FontAwesomeIcon icon={loan.returned ? faCheckCircle : faTimesCircle} className="status-icon" />
                    {loan.returned ? ' Devolvido' : ' Em Aberto'}
                </span>
            </div>

            <div className="card-actions">
                <button onClick={onEdit} className="edit-button">
                    <FontAwesomeIcon icon={faEdit} className="icon" /> Editar
                </button>
                <button onClick={onDelete} className="delete-button">
                    <FontAwesomeIcon icon={faTrashAlt} className="icon" /> Deletar
                </button>
            </div>
        </div>
    );
};

export default LoanCard;
