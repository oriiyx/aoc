export async function getData(url: string) {

    return await fetch(url, {
        headers: {
            cookie: `session=${process.env.AOC_SESSION}`
        }
    }).then((response) => response.text());
}

/**
 * Splits the data by line breaks.
 * @param {string} data
 */
export const lines = (data: string) => data.split('\n');

/**
 * Splits the data by line breaks and parses string values into numbers.
 * @param {string} data
 */
export const numbers = (data: string) => lines(data).map((value) => +value);

/**
 * Splits the data by a custom separator.
 * @param {string | RegExp} separator - the custom separator
 * @returns { (data: string) => string[] }
 */
export const splitBy = (separator: string | RegExp) => (data: string) =>
  data.split(separator);

export const isNumeric = (string: string) => Number.isFinite(+string)
