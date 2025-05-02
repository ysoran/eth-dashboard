"use client"

import React, { useEffect, useState } from "react";
import DashboardSidebar from "../components/DashboardSidebar";
import DashboardHeader from "../components/DashboardHeader";
import BlockchainCard from "../components/BlockchainCard";
import { TopAddresses } from "../components/graphs/TopAddresses";
import { TransactionsPerBlock } from "../components/graphs/TransactionsPerBlock";
import { TxStatusPie } from "../components/graphs/TxStatusPie";
import { ValueDistribution } from "../components/graphs/ValueDistribution";
import { VolumeOverTime } from "../components/graphs/VolumeOverTime";
import TransactionTable from "../components/TransactionTable";

interface Transaction {
  hash: string;
  block: string;
  amount: string;
  status: string;
  from: string;
  to: string;
}

const DashboardPage = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/api/transactions")
      .then((res) => res.json())
      .then((data) => {
        const formatted = data.map((txn: any) => ({
          hash: txn.hash,
          block: txn.blockNumber,
          amount: txn.value,
          status: txn.status,
          from: txn.from,
          to: txn.to,
        }));
        setTransactions(formatted);
      })
      .catch((err) => console.error("Fetch error:", err));

    const ws = new WebSocket("ws://localhost:8080/ws/status");
    ws.onmessage = (event) => {
      const updated = JSON.parse(event.data);
      setTransactions((prev) =>
        prev.map((txn) =>
          txn.hash === updated.hash ? updated : txn
        )
      );
    };
    ws.onerror = (err) => console.error("WebSocket error:", err);
    return () => ws.close();
  }, []);

  return (
    <div className="flex min-h-screen">
      <DashboardSidebar />
      <main className="flex-1 p-6 bg-gray-100 dark:bg-black">
        <DashboardHeader />

        {/* Metric Cards */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 my-6">
          <BlockchainCard title="Total Transactions" value="2,345" />
          <BlockchainCard title="Total ETH Volume" value="10,000 ETH" />
          <BlockchainCard title="Pending Transactions" value="123" />
        </div>

        {/* Graphs */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <TopAddresses />
          <TransactionsPerBlock />
          <TxStatusPie />
          <ValueDistribution />
          <VolumeOverTime />
        </div>

      
      </main>
    </div>
  );
};

export default DashboardPage;
