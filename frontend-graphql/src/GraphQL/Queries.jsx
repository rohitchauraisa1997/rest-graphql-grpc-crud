import {gql} from '@apollo/client'

export const LOAD_VIDEOS = gql`
    query findVideos{
        videos{
        id
        title
        url
        author{
            id
            name
        }
        }
    }
`