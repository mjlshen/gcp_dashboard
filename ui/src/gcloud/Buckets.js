import React, { useState, useEffect } from 'react';
import { DataGrid } from '@material-ui/data-grid'
import axios from 'axios'

export default function Buckets() {
  const [buckets, setBuckets] = useState([]);
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
      width: 300,
    },
    {
      field: 'location',
      headerName: 'Location',
      width: 150,
    },
    {
      field: 'storageClass',
      headerName: 'Storage Class',
      width: 125,
    },
    {
      field: 'versioningEnabled',
      headerName: 'Versioning Enabled',
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
      axios.get("http://localhost:8082/api/v1/gcloud/buckets")
          .then(res => {
              setBuckets(res.data)
          })
  }, []);

  return (
      <div style={{ height: '100vh', width: '100%' }}>
          <div style={{ display: 'flex', height: '100%' }}>
            <div style={{ flexGrow: 1 }}>
                <DataGrid autoPageSize rows={buckets} columns={columns} />
              </div>
          </div>
      </div>
  );
}