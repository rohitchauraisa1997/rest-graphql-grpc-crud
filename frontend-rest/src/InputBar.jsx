import Button from '@mui/material/Button';
import TextField from "@mui/material/TextField";
import {Typography} from "@mui/material";
import { useState } from 'react';

function InputBar(props){

    const [videoTitle, setVideoTitle] = useState("")
    const [url, setUrl] = useState("")
    const [authorName, setAuthorName] = useState("")

    const handleVideoAdd = () =>{
        {
            fetch("http://localhost:3001/video/add",
                {
                    method:"POST",
                    body: JSON.stringify({
                        "title": videoTitle,
                        "url":url,
                        "author":{
                            //TODO: implement user logic.
                            "id": "",
                            "name":authorName
                        }
                    }),
                    headers:{
                        "Content-Type":"application/json"
                    }
                }
            ).then(response => {
                if (response.ok) {
                return response.json();
                } else {
                throw new Error(`Network response was not ok (${response.status})`);
                }
            })
            .then(data => {
                // Use the fetched data
                // Create a new array with the updated data
                const responseObject = {
                    "id": data.id,
                    "title": data.title,
                    "url":data.url,
                    "author": data.author
                }
                // Initialize the updatedResponseObjects with responseObject
                let updatedResponseObjects = [responseObject];

                // If props.videoDetailRows is not empty, append the new object to it
                if (props.videoDetailRows && props.videoDetailRows.length > 0) {
                    updatedResponseObjects = [...props.videoDetailRows, responseObject];
                }
                // Update the state using setVideoDetailRows with the new array
                props.setVideoDetailRows(updatedResponseObjects);
                setVideoTitle("")
                setUrl("");
                setAuthorName("");
            })
            .catch(error => {
                // Handle any errors that occurred during the fetch
                console.error('Fetch error:', error);
                window.alert(error.message);
            })
        }
    }

    return (
        <div style={{marginTop:50, marginBottom:50}}>
            <div style={{"display":"flex",justifyContent:"center"}}>
                <Typography variant={"h6"}>
                    Video Library [REST Server]
                </Typography>
            </div>
            <div style={{display:"flex", justifyContent:"center",marginTop:50, marginBottom:50}}>
            <TextField
                    id="video-title"
                    label="Video Title"
                    variant="outlined"
                    value={videoTitle}
                    style={{marginRight:20}}
                    onChange={(evant) => {
                        let elemt = evant.target;
                        setVideoTitle(elemt.value);
                    }}
                />
                <TextField
                    id="video-url"
                    label="Video Url"
                    variant="outlined"
                    value={url}
                    style={{marginRight:20}}
                    onChange={(evant) => {
                        let elemt = evant.target;
                        setUrl(elemt.value);
                    }}
                />
                <TextField
                    id="video-author"
                    label="Video Author"
                    variant="outlined"
                    value={authorName}
                    style={{marginRight:40}}
                    onChange={(evant) => {
                        let elemt = evant.target;
                        setAuthorName(elemt.value);
                    }}
                />
                <Button 
                size={"large"} 
                variant="contained"
                onClick={handleVideoAdd}
                >
                    Register Video
                </Button>
            </div>
        </div>
    )
}

export default InputBar