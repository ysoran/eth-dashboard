import Link from 'next/link';

const DashboardSidebar = () => (
  <aside className="w-64 bg-gray-100 dark:bg-[#121212] p-4 h-screen border-r border-gray-300 dark:border-gray-700">
    <ul className="space-y-2">
      <li>
        <Link href="/dashboard" className="text-gray-800 dark:text-gray-200 hover:underline">
          Dashboard
        </Link>
      </li>
      <li>
        <Link href="/transactions" className="text-gray-800 dark:text-gray-200 hover:underline">
          Transactions
        </Link>
      </li>
    </ul>
  </aside>
);

export default DashboardSidebar;