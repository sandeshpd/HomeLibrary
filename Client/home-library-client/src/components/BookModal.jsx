import { useEffect, useState } from "react";


function BookModal({ onClose, mode = "add", bookData = {}, onSubmit }) {
  const [formData, setFormData] = useState({
    book_id: "",
    title: "",
    author: "",
    price: "",
    cover_path: "",
    language: "",
    publication: "",
    genre: "",
  });

  useEffect(() => {
    if (mode === "edit" && bookData) {
      setFormData({
        book_id: bookData.book_id || "",
        title: bookData.title || "",
        author: bookData.author || "",
        price: bookData.price || "",
        cover_path: bookData.cover_path || "",
        language: bookData.language || "",
        publication: bookData.publication || "",
        genre: bookData.genre ? bookData.genre.map(g => g.genre_name).join(", ") : "",
      });
    }
  }, [mode, bookData]);

  // Handle form input changes
  const handleInputChange = (e) => {
    const { id, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [id]: value,
    }));
  };

  // Handle form submission
  const handleSubmitBookForm = (e) => {
    e.preventDefault();
    onSubmit(formData, mode);
    console.log("Book data would be sent to the backend.")
    onClose();
  };

  return (
    <>
      <div className="fixed inset-0 flex items-center justify-center z-50">
        <div className="absolute inset-0 bg-black opacity-50" onClick={onClose}></div>
        <div className="bg-[#e6cfa7] rounded-lg shadow-lg p-6 relative z-10 w-[20rem] sm:min-w-[35rem]">
          <button
            className="absolute top-0 right-2 text-gray-500 hover:text-gray-700 cursor-pointer text-3xl"
            onClick={onClose}
          >
            &times;
          </button>
          <h1 className="text-center my-2">
            <span className="text-2xl">
              {mode === "add" ? "Add Book..." : "Edit book..."}
            </span>
          </h1>
          <div>
            <form
              className="flex flex-col gap-4"
              onSubmit={handleSubmitBookForm}
            >
              <input
                type="text"
                id="book_id"
                placeholder="Book ID"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                disabled={mode === "edit"}
                value={formData.book_id}
              />
              <input
                type="text"
                id="title"
                placeholder="Title"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.title}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="author"
                placeholder="Author"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.author}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="price"
                placeholder="Price"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.price}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="cover_path"
                placeholder="Cover"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.cover_path}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="language"
                placeholder="Language"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.language}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="publication"
                placeholder="Publication"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.publication}
                onChange={handleInputChange}
              />
              <input
                type="text"
                id="genre"
                placeholder="Genre"
                className="border border-slate-700 p-2 rounded-md focus:outline-1"
                value={formData.genre}
                onChange={handleInputChange}
              />
              <div className="flex justify-end gap-4">
                {mode === "add" && (
                  <button
                    className="border-2 border-[#3a4b41] bg-inherit text-[#3a4b41] p-2 rounded-md cursor-pointer hover:opacity-95 font-semibold"
                  >
                    Submit & Next
                  </button>
                )}
                <button
                  className="bg-[#3a4b41] text-[#e6cfa7] font-semibold p-2 rounded-md cursor-pointer hover:opacity-95"
                >
                  {mode === "add" ? "Submit" : "Save changes"}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  )
};

export default BookModal;