import React from 'react';
import { Box, Button, OutlinedInput, Stack, TextField, Typography } from "@mui/material";
import { X } from 'react-feather';
import ProrateButton from "./ProrateButton";
import {Prorate} from "./ResultsList";

export type Investor = {
    name?: string | null;
    requestedAmount?: number | null;
    averageAmount?: number | null;
};

type AllocationFormProps = {
    setData: (response: Prorate[]) => void;
};

const AllocationForm = ({ setData }: AllocationFormProps) => {
    const [allocationAmount, setAllocationAmount] = React.useState<number>(0);
    const [investors, setInvestors] = React.useState<Investor[]>([{
        name: '',
        requestedAmount: null,
        averageAmount: null,
    }])

    const addInvestor = () => {
        setInvestors([...investors, {
            name: '',
            requestedAmount: null,
            averageAmount: null,
        }])
    }

    const handleAllocationAmountChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
        setAllocationAmount(Number(e.target.value));
    }

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>, index: number) => {
        const { name, value } = e.target;
        setInvestors(prevState => {
            if (name === 'name') {
                (prevState[index] as any)[name] = value;
            } else {
                (prevState[index] as any)[name] = Number(value);
            }
            return [...prevState];
        })
    }

    const handleRemoveInvestor = (index: number) => {
        const currInvestors = [...investors];
        currInvestors.splice(index, 1);
        setInvestors(currInvestors);
    }

    return (
        <Box component={'form'} border={'1px solid grey'} borderRadius={'8px'} height={'75vh'}>
            <Stack spacing={2} margin={'10px'} height={'100%'} display={'flex'} justifyContent={'space-between'}>
                <Stack spacing={2}>
                    <Typography fontSize={'18px'} variant={'h6'}>Total Available Allocation</Typography>
                    <OutlinedInput
                        onChange={handleAllocationAmountChange}
                        name={'allocationAmount'}
                        required
                        placeholder={' Allocation'}
                        sx={{ width: '30%'}}
                        startAdornment={'$'}
                        type={'number'}
                    />
                </Stack>
                <Stack spacing={2}>
                    <Typography fontSize={'18px'} variant={'h6'}>Investor Breakdown</Typography>
                    {investors.map((investor, index) => {
                        return (
                            <Stack spacing={2} key={index}>
                                <Stack direction={'row'} spacing={2}>
                                    <TextField
                                        onChange={(e) => handleChange(e, index)}
                                        name={'name'}
                                        required
                                        placeholder={'Name'}
                                        value={investor.name}
                                    />
                                    <OutlinedInput
                                        required
                                        name={'requestedAmount'}
                                        onChange={(e) => handleChange(e, index)}
                                        placeholder={' Requested Amount'}
                                        sx={{ width: '30%'}}
                                        startAdornment={'$'}
                                        type={'number'}
                                        value={investor.requestedAmount}
                                    />
                                    <OutlinedInput
                                        required
                                        name={'averageAmount'}
                                        onChange={(e) => handleChange(e, index)}
                                        placeholder={' Average Amount'}
                                        sx={{ width: '30%'}}
                                        startAdornment={'$'}
                                        type={'number'}
                                        value={investor.averageAmount}
                                    />
                                    {investors.length > 1 && (
                                        <Button onClick={() => handleRemoveInvestor(index)}><X /></Button>
                                    )}
                                </Stack>
                            </Stack>
                        )
                    })}
                    {investors.length > 1 && investors.length < 4 && (
                        <Box display={'flex'} justifyContent={'center'} width={'100%'}>
                            <Button onClick={addInvestor} sx={{ border: '1px solid black', width: '50%' }}>Add Investor</Button>
                        </Box>
                    )}
                </Stack>
                <Box alignItems={'flex-end'} display={'flex'} >
                    <ProrateButton allocationAmount={allocationAmount} investorAmounts={investors} setData={setData} />
                </Box>
            </Stack>
        </Box>
    );
};

export default AllocationForm;