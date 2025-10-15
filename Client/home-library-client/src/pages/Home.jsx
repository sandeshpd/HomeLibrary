import { useState } from "react";
import BookModal from "../components/BookModal";
import { Link } from "react-router-dom";

function Home() {
  const [openAddBookModal, setOpenAddBookModal] = useState(false);

  const handleOpenModalBox = () => {
    setOpenAddBookModal(true);
    console.log("Pressed add book to open modal", openAddBookModal);
  };

  const handleCloseModal = () => {
    setOpenAddBookModal(false);
    console.log("pressed Close modal icon", openAddBookModal);
  }

  const handleSubmit = (formData, mode) => {
    if (mode === "add") {
      console.log("Adding book:", formData);
      // TODO: Send formData to the backend for adding a book
    }
  };
  return (
    <>
      <div className='flex max-w-lg mx-auto items-center h-[100vh]'>
        <div className="flex flex-col justify-center w-full gap-9 mx-[2rem] sm:flex-row">
          <button
            className="border-2 border-[#3a4b41] cursor-pointer hover:opacity-90 p-3 
                        rounded-lg shadow-lg text-[#3a4b41] font-semibold"
            onClick={handleOpenModalBox}
          >
            Add Book
          </button>
          <button
            className="bg-[#3a4b41] cursor-pointer hover:opacity-90 p-3 rounded-lg 
                        text-[#e6cfa7] shadow-md shadow-[#495850] font-semibold"
          >
            <Link to={"/list"}>
              View your Collection
            </Link>
          </button>
        </div>
        {openAddBookModal && (
          <>
            <BookModal
              onClose={handleCloseModal}
              mode="add"
              onSubmit={handleSubmit}
            />
          </>
        )}
      </div>
    </>
  )
}

export default Home;