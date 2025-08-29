import { useState } from 'react'
import './App.css'

function App() {
  const [file, setFile] = useState<File | undefined>();
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const file = e.target.files[0];
      setFile(file)
    }
  }
  const handleUpload = async () => {
    if (file) {
      const CHUNK = 5 * 1024 * 1024;
      const uploadId = crypto.randomUUID()
      let START = 0;
      let part = 1;
      while (START < file.size) {
        const FORM = new FormData()
        const slice = file.slice(START, START + CHUNK)
        FORM.append("uploadId", uploadId)
        FORM.append("part", part.toString())
        FORM.append("chunk", slice)
        FORM.append("end", (START + CHUNK >= file.size) ? "True" : "False")
        FORM.forEach((val) => {
          console.log(val)
        })
        await fetch("http://localhost:5000/upload-video", {
          method: "POST",
          body: FORM
        })
        START += CHUNK
        part++
      }
    }
  }

  return (
    <>
      <input type="file" name="upload" accept="video/*" onChange={handleFileChange} />
      <button onClick={handleUpload}>Upload Video</button>
    </>
  )
}

export default App
