import InputBar from "./InputBar"
import VideoTable  from "./VideoTable";
import './App.css';
import { useEffect, useState } from "react";

function App() {
  const [videoDetailRows, setVideoDetailRows] = useState([]);
  useEffect(()=>{
      fetch("http://localhost:8081/videos/all",
      {method:"GET"}).then(
          response=> response.json()).then(
          data=>{
            if (data.videos){
            console.log(data.videos);
            setVideoDetailRows(data.videos)
            }
        }
      )
  },[])

  if (videoDetailRows && videoDetailRows.length>0){ 
    return (
      <div className="app-container"> {/* Wrap the entire app */}
        <InputBar videoDetailRows={videoDetailRows} setVideoDetailRows={setVideoDetailRows}></InputBar>
        <VideoTable videoDetailRows={videoDetailRows} setVideoDetailRows={setVideoDetailRows}></VideoTable>
      </div>  
    );

  }else{
    return(
      <div className="app-container"> {/* Wrap the entire app */}
      <InputBar videoDetailRows={videoDetailRows} setVideoDetailRows={setVideoDetailRows}></InputBar>
      </div>  
    )
  }

}

export default App