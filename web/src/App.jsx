import { useState } from "react";
import axios from "axios";
import { useCopyToClipboard } from "usehooks-ts";

function App() {
  const [formData, setFormData] = useState();
  const [shortenedUrl, setShortenedUrl] = useState();

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.id]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    axios.post("/shorten", formData).then(
      (response) => {
        console.log(response.status, response.data);
        setShortenedUrl(response.data.shortenedUrl);
      },
      (err) => console.log(err)
    );
  };

  return (
    <div className="relative flex h-[100vh]">
      <div className="absolute inset-0 bg-[url('./assets/wallhaven-bg.jpg')] bg-no-repeat bg-cover opacity-60"></div>
      <div className="absolute inset-0 bg-blue-500 opacity-10"></div>
      <div className="mx-auto mt-36 text-center relative z-10 text-slate-900">
        <h1 className="text-3xl font-bold">Paste the URL to be shortened</h1>
        <p className="text-slate-700">
          URL shortener allows to create a shortened link making it easy to
          share
        </p>
        <form
          method="post"
          className="flex items-center space-x-2 mt-6"
          onSubmit={handleSubmit}
        >
          <input
            className="py-2 px-4 rounded-lg h-12 flex-grow"
            type="url"
            id="url"
            placeholder="Enter a link"
            onChange={handleChange}
            required
          />
          <input
            className="h-12 px-6 cursor-pointer bg-gradient-to-tr from-blue-700 to-blue-400 text-white font-bold text-xs uppercase rounded-lg shadow-md shadow-gray-900/10 hover:shadow-lg hover:shadow-gray-900/20 active:opacity-[0.85] whitespace-nowrap"
            type="submit"
            value="Shorten URL"
          />
        </form>
        <div className="mt-12">
          {shortenedUrl ? (
            <>
              <h1 className="text-3xl font-bold">Your shortened URL</h1>
              <ClipboardCopyInput inputValue={shortenedUrl} className="mt-6" />
            </>
          ) : null}
        </div>
      </div>
    </div>
  );
}

const ClipboardCopyInput = ({ inputValue, className }) => {
  // eslint-disable-next-line no-unused-vars
  const [value, copy] = useCopyToClipboard();
  const [copied, setCopied] = useState(false);

  return (
    <div className={`${className} flex items-center space-x-2`}>
      <input
        value={inputValue}
        type="url"
        className="py-2 px-4 rounded-lg h-12 flex-grow backdrop-filter backdrop-blur-lg bg-opacity-30"
        disabled
        readOnly
      />
      <button
        onMouseLeave={() => setCopied(false)}
        onClick={() => {
          copy(inputValue);
          setCopied(true);
        }}
        className="flex items-center h-12 px-6 cursor-pointer bg-gradient-to-tr from-blue-700 to-blue-400 text-white font-bold text-xs uppercase rounded-lg shadow-md shadow-gray-900/10 hover:shadow-lg hover:shadow-gray-900/20 active:opacity-[0.85] whitespace-nowrap"
      >
        {copied ? (
          <>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="size-6 mr-2"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="m4.5 12.75 6 6 9-13.5"
              />
            </svg>
            Copied
          </>
        ) : (
          <>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="size-6 mr-2"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M15.75 17.25v3.375c0 .621-.504 1.125-1.125 1.125h-9.75a1.125 1.125 0 0 1-1.125-1.125V7.875c0-.621.504-1.125 1.125-1.125H6.75a9.06 9.06 0 0 1 1.5.124m7.5 10.376h3.375c.621 0 1.125-.504 1.125-1.125V11.25c0-4.46-3.243-8.161-7.5-8.876a9.06 9.06 0 0 0-1.5-.124H9.375c-.621 0-1.125.504-1.125 1.125v3.5m7.5 10.375H9.375a1.125 1.125 0 0 1-1.125-1.125v-9.25m12 6.625v-1.875a3.375 3.375 0 0 0-3.375-3.375h-1.5a1.125 1.125 0 0 1-1.125-1.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H9.75"
              />
            </svg>
            Copy
          </>
        )}
      </button>
    </div>
  );
};

export default App;
