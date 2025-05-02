"use client"
import React, { useEffect, useState } from "react";
import BlockchainCard from "../components/BlockchainCard";
import DashboardHeader from "../components/DashboardHeader";
import DashboardSidebar from "../components/DashboardSidebar";
import TransactionTable from "../components/TransactionTable";

interface Transaction {
  hash: string;
  block: string;
  amount: string; // still a string in hex
  status: string;
  from: string;
  to: string;
}

const TransactionsPage = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null); // WebSocket state to manage connection

  useEffect(() => {
    // Fetch transactions on initial load
    console.log("Fetching transactions...");
    fetch('http://localhost:8080/api/transactions')
      .then((response) => response.json())
      .then((data) => {
        const formattedData = data.map((txn: any) => ({
          hash: txn.hash,
          block: txn.blockNumber,
          amount: txn.value,
          status: txn.status,
          from: txn.from,
          to: txn.to,
        }));

        console.log("Formatted Transactions:", formattedData); // Log the formatted data
        setTransactions(formattedData);
      })
      .catch((error) => {
        console.error('Error fetching transactions:', error);
      });

    // Set up WebSocket connection for real-time updates
    const socket = new WebSocket('ws://localhost:8080/ws/status');
    socket.onopen = () => {
      console.log('WebSocket connected');
    };

    socket.onmessage = (event) => {
      const updatedTransaction: Transaction = JSON.parse(event.data);
      setTransactions((prevTransactions) =>
        prevTransactions.map((txn) =>
          txn.hash === updatedTransaction.hash ? updatedTransaction : txn
        )
      );
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.onclose = () => {
      console.log('WebSocket disconnected');
    };

    // Clean up WebSocket connection when component unmounts
    return () => {
      socket.close();
    };
  }, []);

  return (
    <div className="flex">
      <DashboardSidebar />
      <main className="flex-1 p-6">
        <DashboardHeader />
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <BlockchainCard title="Total Transactions" value="2,345" />
          <BlockchainCard title="Total Value" value="10,000 ETH" />
          <BlockchainCard title="Pending Transactions" value="123" />
        </div>
        <TransactionTable transactions={transactions} />
      </main>
    </div>
  );
};

export default TransactionsPage;
