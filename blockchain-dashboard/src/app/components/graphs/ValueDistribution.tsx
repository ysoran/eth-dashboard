import React, { useEffect, useState } from "react";
import { XAxis, Bar, YAxis, BarChart, Tooltip, ResponsiveContainer } from 'recharts';

export function ValueDistribution() {
    const [data, setData] = useState([]);
    useEffect(() => {
      fetch("http://localhost:8080/api/stats/value-distribution").then(res => res.json()).then(setData);
    }, []);
  
    return (
      <div className="w-full h-80 p-4 bg-white dark:bg-gray-900 rounded-xl shadow">
        <h2 className="text-lg font-semibold mb-2">ETH Value Distribution</h2>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data}>
            <XAxis dataKey="range" />
            <YAxis />
            <Tooltip />
            <Bar dataKey="count" fill="#ffc658" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    );
  }