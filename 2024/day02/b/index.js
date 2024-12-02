import fs from 'fs'

function main() {
    const data = readFileSync("input.txt")
    const dataNumbers = data.map(row => row.map(Number));
    
    let result = 0
    dataNumbers.forEach((report) => {
        if (isSafe(report)) {
            result += 1;
            return;
        }
        
        for (let i = 0; i < report.length; i++) {
            const modifiedReport = [...report.slice(0, i), ...report.slice(i + 1)];
            if (isSafe(modifiedReport)) {
                result += 1;
                return;
            }
        }
    })

    console.log(result)
}

function isSafe(report) {
    if (Math.abs(report[0] - report[1]) == 0 ||  Math.abs(report[0] - report[1]) > 3) return false;
    
    const direction = (report[0] < report[1]) ? 'asc' : 'desc';
    for (let i = 1; i < report.length - 1; i++) {
        if (Math.abs(report[i] - report[i+1]) == 0 ||  Math.abs(report[i] - report[i+1]) > 3) return false;
        if (direction == 'asc') {
            if (!(report[i] < report[i+1])) return false;
        }
        if (direction == 'desc') {
            if (!(report[i] > report[i+1])) return false;
        }
    }
    return true;
}

function readFileSync(filePath) {
    try {
        const data = fs.readFileSync(filePath, 'utf8').split('\n');
        return data.map(line => line.split(" "));
    } catch (error) {
        console.error('Error reading file:', error);
        throw error;
    }
}

main()