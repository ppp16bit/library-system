import React, { useState, useEffect } from 'react';
import './BookForm.css';

const BookForm = ({ bookToEdit, onSubmit, onCancel }) => {
  const [book, setBook] = useState({
    title: '',
    author: '',
    isbn: '',
    available: true
  });
  const [errors, setErrors] = useState({});

  useEffect(() => {
    if (bookToEdit) {
      setBook(bookToEdit);
    } else {
      setBook({ title: '', author: '', isbn: '' });
    }
    setErrors({});
  }, [bookToEdit]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setBook(prevBook => ({ ...prevBook, [name]: value }));
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
    const bookDataToSend = {
        title: book.title,
        author: book.author,
        isbn: book.isbn,
        available: true,
    };
    console.log('Payload do Livro para envio:', bookDataToSend);
    onSubmit(bookDataToSend);
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
