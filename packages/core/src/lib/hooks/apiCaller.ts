export default async function performApiCalls(count: number, origin: string): Promise<void> {
    console.log(count, origin)
    const url = `https://api.example.com/data?origin=${origin}`;

    for (let i = 0; i < count; i++) {
        fetch(url)
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
    }
}