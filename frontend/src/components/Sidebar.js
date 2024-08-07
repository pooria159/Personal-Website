import React from 'react';

function Sidebar() {
  return (
    <div className="w-1/4 bg-gray-800 text-white h-screen p-5">
      <ul>
        <li className="mb-2"><a href="/about">About</a></li>
        <li className="mb-2"><a href="/interests">Interests</a></li>
        <li className="mb-2"><a href="/languages">Languages</a></li>
        <li className="mb-2"><a href="/skills">Skills</a></li>
        <li className="mb-2"><a href="/licenses">Licenses & Certifications</a></li>
        <li className="mb-2"><a href="/education">Education</a></li>
        <li className="mb-2"><a href="/activity">Activity</a></li>
      </ul>
    </div>
  );
}

export default Sidebar;
