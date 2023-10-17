import React, { useEffect, useState } from "react";
import {useQuery, gql}from '@apollo/client'
import {LOAD_VIDEOS} from '../GraphQL/Queries'

function GetVideos(){

    const {error, loading, data} = useQuery(LOAD_VIDEOS)
    const [allVideos, setAllVideos] = useState([])


    useEffect(()=>{
        console.log(data);
        // console.log(data["videos"]);
        // console.log(data["videos"][0]["title"]);
        setAllVideos(data ? data.videos : []);
    },[data])


    return(
        <>
        {allVideos? (
            allVideos.map((val)=>{
                return <h1 key={val.id}>{val.title}</h1>
            })
        ):(
            <p>Loading..</p>
        )}
        </>
    )

}
export default GetVideos;