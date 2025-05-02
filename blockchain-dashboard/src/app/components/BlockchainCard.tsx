import React from 'react';

interface BlockchainCardProps {
    title: string;
    value: string;
  }
  
  const BlockchainCard: React.FC<BlockchainCardProps> = ({ title, value }) => (
    <div className="bg-white dark:bg-[#1a1a1a] p-6 rounded-2xl shadow-md border border-gray-200 dark:border-gray-700">
      <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-100">{title}</h2>
      <p className="text-2xl font-bold text-blue-600 dark:text-blue-400">{value}</p>
    </div>

  );
  
  export default BlockchainCard;
  