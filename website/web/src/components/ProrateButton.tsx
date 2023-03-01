import React from 'react';
import { Investor } from "./AllocationForm";
import { Button } from "@mui/material";
import {Prorate} from "./ResultsList";

export const ENDPOINT = 'http://localhost:4000';

// @ts-ignore
async function sendRequest(allocationAmount, investorAmounts) {
    let response =  await fetch(`${ENDPOINT}/api/invest`, {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "allocation_amount": allocationAmount,
            "investor_amounts": investorAmounts.map((amount: Investor) => {
                return {
                    "name": amount.name,
                    "requested_amount": amount.requestedAmount,
                    "average_amount": amount.averageAmount,
                };
            }),
        }),
    });

    return await response.json()
}

type ProrateButtonProps = {
    allocationAmount: number;
    investorAmounts: Investor[];
    setData: (response: Prorate[]) => void;
}

const ProrateButton = ({ allocationAmount, investorAmounts, setData }: ProrateButtonProps) => {
    const handleClick = async () => {
        let data = await sendRequest(allocationAmount, investorAmounts);
        setData(data);
    };

    return (
        <>
            <Button
                onClick={handleClick}
                sx={{ border: '1px solid black', borderRadius: '8px', marginBottom: '20px' }}
            >
                Prorate
            </Button>
        </>
    );
};

export default ProrateButton;