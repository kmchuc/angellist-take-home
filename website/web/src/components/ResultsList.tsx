import React from 'react';
import { Stack, Typography } from "@mui/material";

export type Prorate = {
    name: string,
    prorate: number,
}

type ResultsListProps = {
    data?: any;
};

const ResultsList = ({ data = {} }: ResultsListProps) => {
    const result = Object.keys(data).map((el, index) => {
        return (
                <li key={index}>
                    <Typography variant={'h6'}>
                        {el} - ${data[el]}
                    </Typography>
                </li>
        );
    })

    return (
        <Stack direction={'column'} spacing={2} border={'1px solid black'} borderRadius={'8px'} height={'75vh'}>
            <Stack height={'100%'} justifyContent={'center'}>
                <ul style={{ listStyle: "none" }}>
                    {result}
                </ul>
            </Stack>
        </Stack>
    );
};

export default ResultsList;