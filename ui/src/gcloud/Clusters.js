import React, { useState, useEffect } from 'react';
import { DataGrid } from '@material-ui/data-grid'
import axios from 'axios'

export default function Clusters() {
  const [clusters, setClusters] = useState([]);
  const columns = [
    {
      field: 'id',
      hide: true
    },
    {
      field: 'projectId',
      headerName: 'Project ID',
      width: 250,
    },
    {
      field: 'name',
      headerName: 'Name',
      width: 200,
    },
    {
      field: 'location',
      headerName: 'Location',
      width: 125,
    },
    {
      field: 'currentMasterVersion',
      headerName: 'Master Version',
      width: 150,
    },
    {
      field: 'currentNodeVersion',
      headerName: 'Node Version',
      width: 150,
    },
    // {
    //   field: 'projectId',
    //   headerName: 'Console',
    //   width: 150,
    //   renderCell: (params) => (
    //     <Button
    //       variant="contained"
    //       color="primary"
    //       size="small"
    //       endIcon={<ComputerIcon />}
    //       href={"https://console.cloud.google.com/home/dashboard?project=" + params.getValue('projectId')}
    //     >
    //         Console
    //     </Button>
    //   ),
    // },
  ];

  useEffect(() => {
      axios.get("http://localhost:8082/api/v1/gcloud/clusters")
          .then(res => {
              setClusters(res.data)
          })
  }, []);

  return (
      <div style={{ height: '100vh', width: '100%' }}>
          <div style={{ display: 'flex', height: '100%' }}>
            <div style={{ flexGrow: 1 }}>
                <DataGrid autoPageSize rows={clusters} columns={columns} />
              </div>
          </div>
      </div>
  );
}