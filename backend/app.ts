import express from "express";
import path from "path";
import bodyParser from "body-parser";
import bcrypt from "bcryptjs";
import { Request, Response } from "express";

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use((req, res, next) => {
  res.header("Access-Control-Allow-Origin", "*");
  res.header(
    "Access-Control-Allow-Headers",
    "Origin, X-Requested-With, Content-Type, Accept"
  );
  if (req.method === "OPTIONS") {
    res.header("Access-Control-Allow-Methods", "PUT, POST, PATCH, DELETE, GET");
    return res.status(200).json({});
  }
  next();
});
app.use(express.static(path.join(__dirname, "../frontend")));

app.get("/", (req, res) => {
  res.sendFile(path.join(__dirname, "../frontend", "index.html"));
});

app.post("/generate", (req: Request, res: Response) => {
  const { length } = req.body;
  const password = generateRandomPassword(length);
  const hash = bcrypt.hashSync(password, bcrypt.genSaltSync(10));
  const match = bcrypt.compareSync(password, hash);

  res.json({ password, hash, match });
});

function generateRandomPassword(length: number) {
  const charset =
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*";
  let password = "";
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * charset.length);
    password += charset[randomIndex];
  }
  return password;
}

const port = 8080;

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
