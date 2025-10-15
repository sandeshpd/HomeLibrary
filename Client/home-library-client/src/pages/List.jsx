import { HiDotsHorizontal } from "react-icons/hi";
import { headers, seed } from "../data/seedData.js";
import { useEffect, useRef, useState } from "react";
import BookModal from "../components/BookModal.jsx";

function List() {
  const [openMenuId, setOpenMenuId] = useState(null);
  const [openEditModal, setOpenEditModal] = useState(false);
  const [selectedBook, setSelectedBook] = useState(null);
  const [bookToDelete, setBookToDelete] = useState(null);
  const [deleteConfirmationAlert, setDeleteConfirmationAlert] = useState(false);
  const menuRef = useRef(null);

  const handleOpenActionsMenu = (book_id) => {
    // TODO: Open Edit and Delete menu
    setOpenMenuId(openMenuId === book_id ? null : book_id)
  };

  const handleEdit = (book_id) => {
    const book = seed.find((b) => b.book_id === book_id)
    setSelectedBook(book);
    setOpenEditModal(true);
    setOpenMenuId(null);
    console.log(selectedBook);
  };

  const handleDeleteButton = (book_id) => {
    setOpenMenuId(null);
    setBookToDelete(book_id);
    setDeleteConfirmationAlert(true);
    // console.log("Delete confirmation box closed without any action.");
  };

  const handleConfirmDelete = () => {
    console.log(`Deletion of Book no. ${bookToDelete} confirmed.`);
    setDeleteConfirmationAlert(false);
    setBookToDelete(null);
    // TODO: Implement delete logic
  };

  const handleCloseDeleteDialogue = () => {
    setDeleteConfirmationAlert(false);
    console.log("Delete confirmation box closed without any action.");
  };

  const handleCloseModal = () => {
    setOpenEditModal(false);
    setSelectedBook(null);
  };

  const handleSubmit = (formData, mode) => {
    if (mode === "edit") {
      console.log("Editing book:", formData);
      // TODO: Implement Book update logic
    }
  };

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (menuRef.current && !menuRef.current.contains(event.target)) {
        setOpenMenuId(null);
      }
    };

    if (openMenuId !== null) {
      document.addEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [openMenuId]);

  // HTML Rendering
  return (
    <>
      <main className="flex justify-center items-center h-[100vh]">
        <div className="relative bg-white min-w-[85%] h-[35rem] rounded-3xl p-5">
          <table className="w-full">
            <thead className="w-full">
              <tr className="flex w-full">
                {headers.map((label, index) =>
                  <th
                    className="flex justify-around p-3 px-0 border-b border-slate-800 bg-gray-200 w-full"
                    key={index}
                  >
                    <h1 className="text-lg font-semibold text-blue-gray-900 opacity-70">
                      {label}
                    </h1>
                  </th>
                )}
              </tr>
            </thead>

            <tbody className="flex flex-col w-full">
              {seed.map(
                ({ book_id, title, author, price, language, publication, genre }) =>
                  <tr key={book_id} className="flex justify-around relative">
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-md font-normal text-gray-900 truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {title}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-md font-normal text-gray-900 truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {author}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-md font-normal text-gray-900 truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {price}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-md font-normal text-gray-900 truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {language}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-md font-normal text-gray-900 truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {publication}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="truncate max-w-xs overflow-hidden whitespace-nowrap">
                        {genre.map(g => g.genre_name).join(", ")}
                      </p>
                    </td>
                    <td className="p-2 flex flex-wrap justify-center border-b border-gray-500 w-full">
                      <p className="text-2xl font-normal text-gray-900 truncate max-w-xs" >
                        <HiDotsHorizontal
                          className="cursor-pointer hover:bg-gray-300 p-1 hover:rounded-[50%]"
                          onClick={() => handleOpenActionsMenu(book_id)}
                        />
                      </p>
                      {openMenuId === book_id && (
                        <div
                          className="absolute top-0 right-[-60px] bg-white border rounded border-gray-300 shadow-2xl z-10 min-w-30"
                          ref={menuRef}
                        >
                          <button
                            className="block px-4 py-2 text-gray-700 hover:bg-gray-100 w-full text-left cursor-pointer"
                            onClick={() => handleEdit(book_id)}
                          >
                            Edit
                          </button>
                          <button
                            className="block px-4 py-2 text-red-500 hover:bg-gray-100 w-full text-left cursor-pointer"
                            onClick={() => handleDeleteButton(book_id)}
                          >
                            Delete
                          </button>
                        </div>
                      )}
                    </td>
                  </tr>
              )}
            </tbody>
          </table>
        </div>
      </main>
      {deleteConfirmationAlert && (
        <div className="fixed inset-0 flex items-center justify-center z-50">
          <div className="absolute inset-0 bg-black opacity-50" onClick={handleCloseDeleteDialogue}></div>
          <div className="bg-[#e6cfa7] rounded-lg shadow-lg p-6 relative z-10 w-[20rem] sm:min-w-[35rem]">
            <h1 className="text-center my-4">
              <span className="text-xl text-[#3a4b41]">
                Are you sure you want to delete selected item?
              </span>
            </h1>
            <div className="flex justify-end gap-3">
              <button
                className="bg-[#e6cfa7] text-[#3a4b41] border border-[#3a4b41] rounded-md p-2  cursor-pointer hover:opacity-95"
                onClick={handleConfirmDelete}
              >
                Confirm
              </button>
              <button
                className="bg-[#3a4b41] text-[#e6cfa7] rounded-md p-2 cursor-pointer hover:opacity-95"
                onClick={handleCloseDeleteDialogue}
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}
      {openEditModal && (
        <BookModal
          onClose={handleCloseModal}
          mode="edit"
          bookData={selectedBook}
          onSubmit={handleSubmit}
        />
      )}
    </>
  )
};

export default List;