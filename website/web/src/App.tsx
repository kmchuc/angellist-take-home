import React from 'react';
import { Container, Grid, Stack, Typography } from "@mui/material";
import AllocationForm, { Investor } from "./components/AllocationForm";
import ResultsList, {Prorate} from "./components/ResultsList";

export type Allocation = {
    allocationAmount: number;
    investors: Investor[];
};

function App() {
  const [data, setData] = React.useState<Prorate[]>([]);

  return (
    <Container disableGutters maxWidth={'xl'}>
        <Grid
            container
            justifyContent={'center'}
            alignItems={'center'}
            spacing={2}
            sx={{ height: '100vh' }}
        >
            <Grid item xs={7}>
                <Stack spacing={2}>
                    <Typography variant={'h5'}>Inputs</Typography>
                    <AllocationForm setData={setData} />
                </Stack>
            </Grid>
            <Grid item xs={3}>
                <Stack spacing={2}>
                    <Typography variant={'h5'}>Results</Typography>
                    {data && <ResultsList data={data} />}
                </Stack>
            </Grid>
        </Grid>
    </Container>
  );
}

export default App;
