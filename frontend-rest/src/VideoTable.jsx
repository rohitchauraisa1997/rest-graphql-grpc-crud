import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import DeleteIcon from '@mui/icons-material/Delete';

function VideoTable(props){
    const videoDetailRows = props.videoDetailRows

    const handleDelete = (id) =>{
        console.log("handleDelete triggered for ", id);
        fetch("http://localhost:3001/video/delete?id="+id,{
                method:"DELETE",
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
                const videoDetailedRowsAfterDeletion = props.videoDetailRows.filter(obj => obj.id !== id);
                props.setVideoDetailRows(videoDetailedRowsAfterDeletion);
            })
            .catch(error => {
                // Handle any errors that occurred during the fetch
                console.error('Fetch error:', error);
                window.alert(error.message);
            })
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
                        key={row.id}
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