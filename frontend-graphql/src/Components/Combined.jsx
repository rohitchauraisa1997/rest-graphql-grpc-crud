import React, { useState, useEffect } from "react";
import InputBar from "./InputBar";
import VideoTable from "./VideoTable";
import { useQuery, gql } from "@apollo/client";
import { LOAD_VIDEOS } from "../GraphQL/Queries";

function Combined(props) {
  const { error, loading, data } = useQuery(LOAD_VIDEOS);
  const [videoDetailRows, setVideoDetailRows] = useState([]);

  useEffect(() => {
    // Check if data is available and not loading
    if (!loading && data) {
      // Set videoDetailRows with data.videos
      setVideoDetailRows(data.videos);
    }
  }, [data, loading]);

  if (videoDetailRows.length>0){ 
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

export default Combined;
