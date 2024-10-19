/* eslint-disable @typescript-eslint/no-misused-promises */
/* eslint-disable @typescript-eslint/no-unsafe-argument */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useState } from "react";

// Modal component
const Modal: React.FC<{ data: any; onClose: () => void }> = ({
  data,
  onClose,
}) => {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-gray-600 bg-opacity-50">
      <div className="rounded-lg bg-white p-6 shadow-lg">
        <h2 className="mb-4 text-center text-xl font-bold">HELP TIPS</h2>
        <p>{data}</p>
        <button
          className="mt-4 rounded bg-blue-500 px-4 py-2 text-white hover:bg-blue-600"
          onClick={onClose}
        >
          Close
        </button>
      </div>
    </div>
  );
};

const GetHelpButton: React.FC = () => {
  const [isModalOpen, setModalOpen] = useState(false); // Modal open state
  const [helpData, setHelpData] = useState<string>(""); // Help data state

  const handleClick = async () => {
    try {
      const response = await fetch("http://localhost:5001/v1/llm/help-text", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          word: "The boy", // Sending the word in the request body
        }),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const data = await response.json();
      setHelpData(data.message); // Assume the response has a message field
      setModalOpen(true); // Open the modal when data is fetched
    } catch (error) {
      console.error("Error fetching help data:", error);
    }
  };

  const handleCloseModal = () => {
    setModalOpen(false); // Close the modal
  };

  return (
    <>
      <button
        className="mx-4 rounded bg-blue-500 px-4 py-2 font-bold text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400"
        onClick={handleClick}
      >
        Get Help
      </button>

      {/* Show modal when isModalOpen is true */}
      {isModalOpen && <Modal data={helpData} onClose={handleCloseModal} />}
    </>
  );
};

export default GetHelpButton;
