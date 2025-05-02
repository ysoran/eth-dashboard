import React from 'react';

interface Transaction {
  hash: string;
  block: string;
  amount: string; // still a string in hex
  status: string;
  from: string;
  to: string;
}

interface TransactionTableProps {
  transactions: Transaction[];
}

// Convert hex wei to ETH
const toEth = (weiHex: string): string => {
  const wei = parseInt(weiHex, 16);
  return (wei / 1e18).toFixed(4); // 4 decimal places
};

const TransactionTable: React.FC<TransactionTableProps> = ({ transactions }) => (
    <div className="overflow-x-auto w-full">
    <table className="min-w-[900px] w-full table-fixed border-collapse">
      <thead>
        <tr className="bg-gray-100 dark:bg-gray-700">
          <th className="w-[280px] px-4 py-2 text-left">Transaction ID</th>
          <th className="w-[120px] px-4 py-2 text-left">Block</th>
          <th className="w-[120px] px-4 py-2 text-left">Amount</th>
          <th className="w-[220px] px-4 py-2 text-left">From</th>
          <th className="w-[220px] px-4 py-2 text-left">To</th>
          <th className="w-[100px] px-4 py-2 text-left">Status</th>
        </tr>
      </thead>
      <tbody>
        {transactions.map((txn) => (
          <tr key={txn.hash} className="hover:bg-gray-50 dark:hover:bg-gray-800 border-t">
            <td className="px-4 py-2 truncate">{txn.hash}</td>
            <td className="px-4 py-2">{txn.block}</td>
            <td className="px-4 py-2">{toEth(txn.amount)} ETH</td>
            <td className="px-4 py-2 truncate">{txn.from}</td>
            <td className="px-4 py-2 truncate">{txn.to}</td>
            <td className="px-4 py-2 capitalize">{txn.status}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>  
);

export default TransactionTable;
function formatEther(value: any): React.ReactNode {
    throw new Error('Function not implemented.');
}

