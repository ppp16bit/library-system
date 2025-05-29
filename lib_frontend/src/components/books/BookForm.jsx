import React, { useState, useEffect } from 'react';
import './BookForm.css';

const BookForm = ({ bookToEdit, onSubmit, onCancel }) => {
  const [book, setBook] = useState({
    id: '',
    title: '',
    author: '',
    isbn: '',
    available: true
  });
  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (bookToEdit) {
      setBook({
        id: bookToEdit.id || '',
        title: bookToEdit.title || '',
        author: bookToEdit.author || '',
        isbn: bookToEdit.isbn || '',
        available: bookToEdit.available
      });
    } else {
      setBook({
        id: '',
        title: '',
        author: '',
        isbn: '',
        available: true
      });
    }
    setErrors({});
  }, [bookToEdit]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setBook(prevBook => ({
      ...prevBook,
      [name]: type === 'checkbox' ? checked : value
    }));
    setErrors(prevErrors => ({ ...prevErrors, [name]: '' }));
  };

  const validate = () => {
    const newErrors = {};
    if (!book.title.trim()) newErrors.title = 'Título é obrigatório.';
    if (!book.author.trim()) newErrors.author = 'Autor é obrigatório.';
    if (!book.isbn.trim()) newErrors.isbn = 'ISBN é obrigatório.';
    else if (!/^\d{10}(\d{3})?$/.test(book.isbn)) {
      newErrors.isbn = 'ISBN deve ter 10 ou 13 dígitos numéricos.';
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (validate()) {
      onSubmit(book);
    }
  };

  return (
    <div className="book-form-container">
      <h2>{bookToEdit ? 'Editar Livro' : 'Adicionar Novo Livro'}</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="title">Título:</label>
          <input
            type="text"
            id="title"
            name="title"
            value={book.title}
            onChange={handleChange}
            placeholder="Digite o título do livro"
            className={errors.title ? 'input-error' : ''}
          />
          {errors.title && <p className="error-text">{errors.title}</p>}
        </div>
        <div className="form-group">
          <label htmlFor="author">Autor:</label>
          <input
            type="text"
            id="author"
            name="author"
            value={book.author}
            onChange={handleChange}
            placeholder="Digite o nome do autor"
            className={errors.author ? 'input-error' : ''}
          />
          {errors.author && <p className="error-text">{errors.author}</p>}
        </div>
        <div className="form-group">
          <label htmlFor="isbn">ISBN:</label>
          <input
            type="text"
            id="isbn"
            name="isbn"
            value={book.isbn}
            onChange={handleChange}
            placeholder="Digite o ISBN do livro"
            className={errors.isbn ? 'input-error' : ''}
          />
          {errors.isbn && <p className="error-text">{errors.isbn}</p>}
        </div>

        {bookToEdit && (
          <div className="form-group form-checkbox-group">
            <input
              type="checkbox"
              id="available"
              name="available"
              checked={book.available}
              onChange={handleChange}
            />
            <label htmlFor="available">Disponível para Empréstimo?</label>
          </div>
        )}

        <div className="form-actions">
          <button type="submit" className="submit-button">
            {bookToEdit ? 'Salvar Alterações' : 'Adicionar Livro'}
          </button>
          <button type="button" onClick={onCancel} className="cancel-button">
            Cancelar
          </button>
        </div>
      </form>
    </div>
  );
};

export default BookForm;
