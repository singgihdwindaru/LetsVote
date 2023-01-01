using System.Security.Cryptography;
using System.Text;

namespace Vote.Api.Models;

public class Blockchain
{
    public class Block
    {
        public long Timestamp { get; set; }
        public byte[] PreviousBlockHash { get; set; }
        public byte[] CurrentBlockHash { get; set; }
        public byte[] AllData { get; set; }
    }

    public Blockchain(string data, byte[] prevBlockHash)
    {

    }

    public Block NewBlock(string data, byte[] prevBlockHash)
    {
        Block block = new();
        block.Timestamp = new DateTimeOffset(DateTime.UtcNow).ToUnixTimeSeconds();
        block.PreviousBlockHash = prevBlockHash;
        block.CurrentBlockHash = new byte[0];
        block.AllData = System.Text.Encoding.ASCII.GetBytes(data);
        setHash(block);
        return block;
    }
    private void setHash(Block block)
    {
        var sep = System.Text.Encoding.ASCII.GetBytes("");
        byte[] headers = block.PreviousBlockHash.Concat(sep).Concat(block.AllData).ToArray();
        // Create a SHA256   

        byte[] SHA256Result;
        StringBuilder stringBuilder = new StringBuilder();
        using (SHA256 sha256Hash = SHA256.Create())
        {
            // ComputeHash - returns byte array  
            SHA256Result = sha256Hash.ComputeHash(headers);
            for (int i = 0; i < SHA256Result.Length; i++)
            {
                stringBuilder.Append(SHA256Result[i].ToString("x2"));
            }
            stringBuilder.ToString();
        }
        block.CurrentBlockHash = System.Text.Encoding.ASCII.GetBytes(stringBuilder.ToString());
    }
    // func (blockchain *blockchain) AddBlock(data string) {
    // 	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
    // 	newBlock := newBlock(data, PreviousBlock.CurrentBlockHash)
    // 	blockchain.Blocks = append(blockchain.Blocks, newBlock)
    // }

}