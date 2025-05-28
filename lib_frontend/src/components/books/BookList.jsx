import React, { useEffect, useState, forwardRef, useImperativeHandle } from 'react';
import axios from 'axios';
import BookCard from './BookCard';
import './BookList.css';
import { API_BASE_URL } from '../../constants';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';

const BookList = forwardRef(({ onEditBook, onDeleteBook }, ref) => {
  const [books, setBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchBooks = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get(`${API_BASE_URL}/books`);
      setBooks(response.data);
    } catch (err) {
      console.error("Erro ao buscar livros:", err);
      setError("Não foi possível carregar os livros. Tente novamente mais tarde.");
    } finally {
      setLoading(false);
    }
  };

  useImperativeHandle(ref, () => ({
    fetchBooks
  }));

  useEffect(() => {
    fetchBooks();
  }, []);

  if (loading) {
    return (
      <div className="loading-container">
        <FontAwesomeIcon icon={faSpinner} spin size="2x" color="#4682b4" />
        <p>Carregando livros...</p>
      </div>
    );
  }

  if (error) return <p className="error-message">{error}</p>;
  if (books.length === 0) return <p className="no-books-message">Nenhum livro cadastrado.</p>;

  return (
    <div className="book-list-container">
      {books.map(book => (
        <BookCard
          key={book.id}
          book={book}
          onEdit={() => onEditBook(book)}
          onDelete={() => onDeleteBook(book.id)}
        />
      ))}
    </div>
  );
});

export default BookList;
