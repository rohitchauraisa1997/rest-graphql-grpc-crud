import {gql} from '@apollo/client'

export const CREATE_VIDEO_MUTATION = gql`
  mutation CreateVideo($input: NewVideo!) {
    createVideo(input: $input) {
      author {
        id
        name
      }
      id
      title
      url
    }
  }
`; 

// export const DELETE_VIDEO_MUTATION = gql`
//   mutation DeleteVideo($id: ID!) {
//     deleteVideo(id: $id) {
//       title
//       url
//     }
//   }
// `


export const DELETE_VIDEO_MUTATION = gql`
  mutation DeleteVideo($id: String!) {
    deleteVideo(id: $id) {
      id
      title
      url
      author {
        id
        name
      }
    }
  }
`;