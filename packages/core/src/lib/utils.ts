export function generateId(countryCode: string): string {
  const date = new Date();
  const year = date.getFullYear();
  const month = date.getUTCMonth() + 1;

  return `${countryCode}${year}${month}${getRandomNumbers(14)}`;
}

function getRandomNumbers(length: number): string {
  const primes: number[] = [3, 7, 13, 17, 31, 37, 41, 43, 47, 71, 73];

  const randomMultiplier = primes[Math.floor(Math.random() * primes.length)];
  let randomNumberSequence = Math.round(
    Math.random() *
      Math.pow(10, 6) *
      Math.random() *
      Math.pow(10, 6) *
      Math.random() *
      Math.pow(10, 6) *
      randomMultiplier!
  ).toString();

  if (randomNumberSequence.length < length) {
    randomNumberSequence = randomNumberSequence.padStart(length, '0');
  } else if (randomNumberSequence.length > length) {
    randomNumberSequence = randomNumberSequence.substring(0, length);
  }

  return randomNumberSequence;
}
