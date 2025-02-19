export const createEmptyGrid = (rows = 20, cols = 10): number[][] => {
    return Array(rows).fill(0).map(() => Array(cols).fill(0));
  };