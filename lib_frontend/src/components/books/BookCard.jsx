import React from 'react';
import './BookCard.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEdit, faTrashAlt } from '@fortawesome/free-solid-svg-icons';

const BookCard = ({ book, onEdit, onDelete }) => {
    return (
        <div className="book-card">
            <div className="card-field">
                <label htmlFor={`title-${book.id}`}>Título:</label>
                <input type="text" id={`title-${book.id}`} value={book.title || ''} readOnly />
            </div>
            <div className="card-field">
                <label htmlFor={`author-${book.id}`}>Autor:</label>
                <input type="text" id={`author-${book.id}`} value={book.author || ''} readOnly />
            </div>
            <div className="card-field">
                <label htmlFor={`isbn-${book.id}`}>ISBN:</label>
                <input type="text" id={`isbn-${book.id}`} value={book.isbn || ''} readOnly />
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

export default BookCard;