import React, { useEffect, useState } from "react";
import { Cell, Legend, Pie, PieChart, Tooltip, ResponsiveContainer } from 'recharts';

const COLORS = ['#0088FE', '#00C49F', '#FFBB28', '#FF8042'];


export function TxStatusPie() {
    const [data, setData] = useState([]);
    useEffect(() => {
      fetch("http://localhost:8080/api/stats/status-ratio").then(res => res.json()).then(setData);
    }, []);
  
    return (
      <div className="w-full h-80 p-4 bg-white dark:bg-gray-900 rounded-xl shadow">
        <h2 className="text-lg font-semibold mb-2">Transaction Status Ratio</h2>
        <ResponsiveContainer width="100%" height="100%">
          <PieChart>
            <Pie
              data={data}
              dataKey="count"
              nameKey="status"
              cx="50%"
              cy="50%"
              outerRadius={80}
              label
            >
              {data.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
              ))}
            </Pie>
            <Legend />
            <Tooltip />
          </PieChart>
        </ResponsiveContainer>
      </div>
    );
  }