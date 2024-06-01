export function getWeightedRandom() {
  const cumulativeWeights: number[] = [];
  weights.reduce((acc, weight, index) => {
    const total = acc + weight;
    cumulativeWeights[index] = total;
    return total;
  }, 0);

  const random =
    Math.random() * cumulativeWeights[cumulativeWeights.length - 1]!;

  // Find the index corresponding to the first value lower than random number
  for (let i = 0; i < cumulativeWeights.length; i++) {
    if (random < cumulativeWeights[i]!) {
      return values[i];
    }
  }

  // In case of rounding errors
  return values[Math.floor(Math.random() * values.length)];
}

const values = [2, 3, 4, 5, 6, 7, 8, 9, 10];
const weights = [5, 1, 5, 1, 1, 1, 1, 1, 1];
