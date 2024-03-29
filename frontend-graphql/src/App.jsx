import './App.css'
import { ApolloClient, InMemoryCache, ApolloProvider, HttpLink, from } from '@apollo/client'
import {onError} from '@apollo/client/link/error'
import GetVideos from './Components/GetVideos'
import InputBar from './Components/InputBar'
import VideoTable from './Components/VideoTable'
import Combined from './Components/Combined'

const errorLink = onError(({
  graphqlErrors, networkError
})=>{
  if (graphqlErrors){
    graphqlErrors.map(({message, location, path})=>{
      alert(`Graphql error: ${message}`)
    })
  }
}

)
const link = from([
  errorLink,
  new HttpLink({uri:"http://localhost:8082/query"}),
])

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: link,
})


function App() {

  return (
    <ApolloProvider client={client}>
      {/* <InputBar></InputBar> */}
      {/* <GetVideos/> */}
      {/* <VideoTable></VideoTable> */}
      {<Combined></Combined>}

    </ApolloProvider>
  )
}

export default App
