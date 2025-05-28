import React, { useState, useRef } from 'react';
import './App.css';
import BookList from './components/books/BookList';
import BookForm from './components/books/BookForm';
import UserList from './components/users/UserList';
import UserForm from './components/users/UserForm';
import LoanList from './components/loans/LoanList';
import LoanForm from './components/loans/LoanForm';
import axios from 'axios';
import { API_BASE_URL } from './constants';

import { CSSTransition } from 'react-transition-group';

function App() {
  const [currentView, setCurrentView] = useState('books');
  const [showForm, setShowForm] = useState(false);
  const [bookToEdit, setBookToEdit] = useState(null);
  const [userToEdit, setUserToEdit] = useState(null);
  const [loanToEdit, setLoanToEdit] = useState(null);

  const bookListRef = useRef(null);
  const bookFormRef = useRef(null);
  const userListRef = useRef(null);
  const userFormRef = useRef(null);
  const loanListRef = useRef(null);
  const loanFormRef = useRef(null);

  const handleSaveBook = async (bookData) => {
    try {
      if (bookData.id) {
        await axios.put(`${API_BASE_URL}/books/${bookData.id}`, bookData);
        alert('Livro atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/books`, bookData);
        alert('Livro adicionado com sucesso!');
      }
      setShowForm(false);
      setBookToEdit(null);
      if (bookListRef.current) {
        bookListRef.current.fetchBooks();
      }
    } catch (error) {
      console.error("Erro ao salvar livro:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar livro: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditBook = (book) => {
    setBookToEdit(book);
    setShowForm(true);
  };

  const handleDeleteBook = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este livro?')) {
      try {
        await axios.delete(`${API_BASE_URL}/books/${id}`);
        alert('Livro deletado com sucesso!');
        if (bookListRef.current) {
          bookListRef.current.fetchBooks();
        }
      } catch (error) {
        console.error("Erro ao deletar livro:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar livro: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };

  const handleSaveUser = async (userData) => {
    try {
      if (userData.id) {
        await axios.put(`${API_BASE_URL}/users/${userData.id}`, userData);
        alert('Usuário atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/users`, userData);
        alert('Usuário adicionado com sucesso!');
      }
      setShowForm(false);
      setUserToEdit(null);
      if (userListRef.current) {
        userListRef.current.fetchUsers();
      }
    } catch (error) {
      console.error("Erro ao salvar usuário:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar usuário: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditUser = (user) => {
    setUserToEdit(user);
    setShowForm(true);
  };

  const handleDeleteUser = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este usuário?')) {
      try {
        await axios.delete(`${API_BASE_URL}/users/${id}`);
        alert('Usuário deletado com sucesso!');
        if (userListRef.current) {
          userListRef.current.fetchUsers();
        }
      } catch (error) {
        console.error("Erro ao deletar usuário:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar usuário: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };

  const handleSaveLoan = async (loanData) => {
    try {
      if (loanData.id) {
        await axios.put(`${API_BASE_URL}/loans/${loanData.id}`, loanData);
        alert('Empréstimo atualizado com sucesso!');
      } else {
        await axios.post(`${API_BASE_URL}/loans`, loanData);
        alert('Empréstimo registrado com sucesso!');
      }
      setShowForm(false);
      setLoanToEdit(null);
      if (loanListRef.current) {
        loanListRef.current.fetchLoans();
      }
    } catch (error) {
      console.error("Erro ao salvar empréstimo:", error.response ? error.response.data : error.message);
      alert('Erro ao salvar empréstimo: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
    }
  };

  const handleEditLoan = (loan) => {
    setLoanToEdit(loan);
    setShowForm(true);
  };

  const handleDeleteLoan = async (id) => {
    if (window.confirm('Tem certeza que deseja deletar este empréstimo?')) {
      try {
        await axios.delete(`${API_BASE_URL}/loans/${id}`);
        alert('Empréstimo deletado com sucesso!');
        if (loanListRef.current) {
          loanListRef.current.fetchLoans();
        }
      } catch (error) {
        console.error("Erro ao deletar empréstimo:", error.response ? error.response.data : error.message);
        alert('Erro ao deletar empréstimo: ' + (error.response?.data?.message || 'Verifique o console para mais detalhes.'));
      }
    }
  };


  const handleCancelForm = () => {
    setShowForm(false);
    setBookToEdit(null);
    setUserToEdit(null);
    setLoanToEdit(null);
  };

  const handleAddButton = () => {
    setBookToEdit(null);
    setUserToEdit(null);
    setLoanToEdit(null);
    setShowForm(true);
  };

  return (
    <div className="app-container">
      <h1>Sistema de Biblioteca</h1>

      <div className="navigation-buttons">
        <button
          onClick={() => { setCurrentView('books'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'books' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Livros
        </button>
        <button
          onClick={() => { setCurrentView('users'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'users' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Usuários
        </button>
        <button
          onClick={() => { setCurrentView('loans'); setShowForm(false); handleCancelForm(); }}
          className={currentView === 'loans' ? 'active-nav-button' : 'nav-button'}
        >
          Gerenciar Empréstimos
        </button>
      </div>

      {!showForm && currentView !== 'loans' && (
        <button onClick={handleAddButton} className="add-item-button">
          Adicionar Novo {currentView === 'books' ? 'Livro' : 'Usuário'}
        </button>
      )}
      {!showForm && currentView === 'loans' && (
        <button onClick={handleAddButton} className="add-item-button">
          Registrar Novo Empréstimo
        </button>
      )}

      <CSSTransition
        in={showForm && currentView === 'books'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={bookFormRef}
      >
        <div ref={bookFormRef}>
          <BookForm
            bookToEdit={bookToEdit}
            onSubmit={handleSaveBook}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={showForm && currentView === 'users'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={userFormRef}
      >
        <div ref={userFormRef}>
          <UserForm
            userToEdit={userToEdit}
            onSubmit={handleSaveUser}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={showForm && currentView === 'loans'} // <<< Transição para LoanForm
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={loanFormRef}
      >
        <div ref={loanFormRef}>
          <LoanForm
            loanToEdit={loanToEdit}
            onSubmit={handleSaveLoan}
            onCancel={handleCancelForm}
          />
        </div>
      </CSSTransition>


      {/* Condicional para renderizar a lista apropriada */}
      <CSSTransition
        in={!showForm && currentView === 'books'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={bookListRef}
      >
        <div ref={bookListRef}>
          <BookList
            onEditBook={handleEditBook}
            onDeleteBook={handleDeleteBook}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={!showForm && currentView === 'users'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={userListRef}
      >
        <div ref={userListRef}>
          <UserList
            onEditUser={handleEditUser}
            onDeleteUser={handleDeleteUser}
          />
        </div>
      </CSSTransition>

      <CSSTransition
        in={!showForm && currentView === 'loans'}
        timeout={300}
        classNames="fade"
        unmountOnExit
        nodeRef={loanListRef}
      >
        <div ref={loanListRef}>
          <LoanList
            onEditLoan={handleEditLoan}
            onDeleteLoan={handleDeleteLoan}
          />
        </div>
      </CSSTransition>
    </div>
  );
}

export default App;