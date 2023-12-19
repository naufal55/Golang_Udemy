const https = require("https");
const http = require("http");
const cheerio = require("cheerio");

async function dataMember(url) {
  return new Promise((resolve, reject) => {
    let data;

    data = url.includes("https") ? https : http;

    data
      .get(url, (res) => {
        let tempData = "";

        res.on("data", (value) => {
          // tampung semua buffer
          tempData += value;
        });

        res.on("end", () => {
          resolve(tempData);
        });
      })
      .on("error", (e) => {
        reject(e);
      });
  });
}

async function main() {
  const grup = 'JKT48'
  const headUrl =
    `http://stage48.net/wiki/index.php?title=Category:${grup}_Songs&pagefrom=`;

  const tailUrl = {
    AKB48 : [
      "",
      "Happy+End",
      `Let%27s+get+"Ato+1+cm"`,
      "Shiroi+Shirts",
    ],
    JKT48 : [
      "",
      "Koko+ni+Ita+Koto+%28JKT48+Song%29"
    ],
  }

  try {
    let final = [];
    const data = await Promise.all(tailUrl[grup.toUpperCase()].map((vl) => dataMember(headUrl+vl)));

    for (const value of data) {
      const $ = cheerio.load(value);
      const jsonData = $("#mw-pages li").append("|").text().split("|");
      final.push(
        ...jsonData.slice(0, jsonData.findIndex((val) => val == "Log in"))
      );
    }

    console.log(final.length);
    console.log(final);
    console.log(final.reverse());

  } catch (error) {
    console.log(error);
  }
}

class User {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }
}

class UserSlice {
  constructor(users) {
    this.users = users;
  }
}

// Contoh penggunaan
const users = [
  new User("Tom", 20),
  new User("Jerry", 18),
  new User("Alice", 30),
];

// Cara benar untuk membuat instance UserSlice
//   const lagi = new UserSlice(users);

// UserSlic
// Menampilkan hasil
// console.log(UserSlice);

main();
