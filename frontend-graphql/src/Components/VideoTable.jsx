import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import DeleteIcon from '@mui/icons-material/Delete';
import {useQuery,useMutation, gql}from '@apollo/client'
import {LOAD_VIDEOS} from '../GraphQL/Queries'
import {DELETE_VIDEO_MUTATION} from '../GraphQL/Mutations'

function VideoTable(props){
    const videoDetailRows = props.videoDetailRows
    const [deleteVideo, {error}] = useMutation(DELETE_VIDEO_MUTATION)
    
    const handleDelete = async(idToDelete) =>{
        try {
            const { data } = await deleteVideo({ variables: { "id": idToDelete } });
            const videoDetailedRowsAfterDeletion = videoDetailRows.filter(obj => obj.id !== idToDelete);
            props.setVideoDetailRows(videoDetailedRowsAfterDeletion);
    }catch(error){
            console.error("Error deleting video: ",error )
            window.alert(error)
        }
    }

    if (videoDetailRows.length>0){ 
        return (
        <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableHead>
                <TableRow>
                <TableCell>Video Title</TableCell>
                <TableCell >Video URL</TableCell>
                <TableCell >Author </TableCell>
                <TableCell >Action </TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {videoDetailRows.map((row,index) => (
                <TableRow
                    key={row.id+index}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">{row.title}</TableCell>
                    <TableCell>
                    <a href={row.title} target="_blank" rel="noopener noreferrer">
                        {row.url}
                    </a>
                    </TableCell>
                    <TableCell>{row.author.name}</TableCell>
                    <TableCell>
                        <Button 
                            variant="outlined" 
                            startIcon={<DeleteIcon />}
                            onClick={() => handleDelete(row.id)}
                        >
                            Delete 
                        </Button>
                    </TableCell>
                </TableRow>
                ))}
            </TableBody>
            </Table>
        </TableContainer>
        );
    }
    
    }
    

export default VideoTable