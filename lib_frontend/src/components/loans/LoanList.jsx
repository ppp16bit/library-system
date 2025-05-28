import React from 'react';
import LoanCard from './LoanCard';
import './LoanList.css';

const LoanList = ({ loans, onEditLoan, onDeleteLoan }) => {
    if (!loans) {
        return (
            <div className="loading-container">
                <p>Carregando empréstimos...</p>
            </div>
        );
    }

    if (loans.length === 0) return <p className="no-items-message">Nenhum empréstimo cadastrado.</p>;

    return (
        <div className="loan-list-container">
            {loans.map(loan => (
                <LoanCard
                    key={loan.id}
                    loan={loan}
                    onEdit={() => onEditLoan(loan)}
                    onDelete={() => onDeleteLoan(loan.id)}
                />
            ))}
        </div>
    );
};

export default LoanList;
