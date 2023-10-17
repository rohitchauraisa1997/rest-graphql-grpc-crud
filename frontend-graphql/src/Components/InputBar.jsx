import Button from '@mui/material/Button';
import TextField from "@mui/material/TextField";
import {Typography} from "@mui/material";
import { useState } from 'react';
import {CREATE_VIDEO_MUTATION} from '../GraphQL/Mutations'
import { useMutation } from '@apollo/client';

function InputBar(props){

    const [videoTitle, setVideoTitle] = useState("")
    const [url, setUrl] = useState("")
    const [authorName, setAuthorName] = useState("")

    const [createVideo, {error}] = useMutation(CREATE_VIDEO_MUTATION)

    const handleVideoAdd= async ()=>{
        try {
            const { data } = await createVideo({
                variables:{
                    input:{
                        title:videoTitle,
                        url:url,
                        author:authorName,
                    }
                }
            })

            const responseObject = {
                "id": data.createVideo.id,
                "title": data.createVideo.title,
                "url":data.createVideo.url,
                "author": data.createVideo.author
            }
            console.log("Created Video:", responseObject);
            let updatedResponseObjects = [responseObject];

            // If props.videoDetailRows is not empty, append the new object to it
            if (props.videoDetailRows.length > 0) {
                updatedResponseObjects = [...props.videoDetailRows, responseObject];
            }
            // Update the state using setVideoDetailRows with the new array
            props.setVideoDetailRows(updatedResponseObjects);
            setVideoTitle("")
            setUrl("");
            setAuthorName("");
        }catch(error){
            console.error("Error creating video: ",error )
            window.alert(error)
        }

    }

    return (
        <div style={{marginTop:50, marginBottom:50}}>
            <div style={{"display":"flex",justifyContent:"center"}}>
                <Typography variant={"h6"}>
                    Video Library [Graphql Server]
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