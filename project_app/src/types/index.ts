
export interface DataRecord {
    id: number
    Age : number
    Work_Class : string
    Education : string
    Marital_Status :  string
    Occupation : string
    Relationship: string
    Race: string
    Sex: string
    Capital_Gain: number
    Capital_Loss: number
    Hours_Per_Week: number
    Native_Country: number
    Income : string
}

export interface DataPagination {
    page: number
    page_size: number
    total_pages: number
    total_records: number
}

