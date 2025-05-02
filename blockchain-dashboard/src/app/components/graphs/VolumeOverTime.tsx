import React, { useEffect, useState } from "react";
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer } from 'recharts';

export function VolumeOverTime() {
    const [data, setData] = useState([]);
    useEffect(() => {
      fetch("http://localhost:8080/api/stats/volume-over-time").then(res => res.json()).then(setData);
    }, []);
  
    return (
      <div className="w-full h-80 p-4 bg-white dark:bg-gray-900 rounded-xl shadow">
        <h2 className="text-lg font-semibold mb-2">Transaction Volume Over Time</h2>
        <ResponsiveContainer width="100%" height="100%">
          <LineChart data={data}>
            <XAxis dataKey="blockNumber" />
            <YAxis />
            <Tooltip />
            <Line type="monotone" dataKey="txCount" stroke="#8884d8" />
          </LineChart>
        </ResponsiveContainer>
      </div>
    );
  }