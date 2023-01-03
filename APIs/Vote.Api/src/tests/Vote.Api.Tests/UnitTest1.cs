using System;
using Vote.Api.Models;
using Xunit.Abstractions;

namespace Vote.Api.Tests;

public class UnitTest1
{
    private readonly ITestOutputHelper output;
    public UnitTest1(ITestOutputHelper output)
    {
        this.output = output;
    }
    private static TestTable[] tcCreateBlock
    {
        get
        {
            return new TestTable[]
            {
                new TestTable {
                    TestName = "#1 Success",
                    Args= new[] {"123"},
                    ExpectedResult=new[] {"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3"},
                    WantError = false,
                },
                new TestTable {
                    TestName = "#2 Success create multiple blocks",
                    Args=new []{"123","321"},
                    ExpectedResult=new []{"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3","829abefbb27be6435d8f57f6bd4b7a327f0e38b04d41d1428f2822da27c29a0b"},
                    WantError = false,
                }
                // ,new TestTable {
                //     TestName = "#3 Error",
                //     Args="",
                //     ExpectedResult="",
                //     WantError = false,
                // }
            };
        }
    }
    public static TheoryData<TestTableBuilder> tdCreateBlock()
    {
        return TestTable.BuildTestTable(tcCreateBlock);
    }

    [Theory]
    [MemberData(nameof(tdCreateBlock))]
    public void TestCreateBlock(TestTableBuilder Case)
    {
        TestTable testData = tcCreateBlock[Case.Index];
        // testData.Mock.Invoke();

        var expectedResult = testData.ExpectedResult as IEnumerable<string>;
        string[]? arg = testData.Args as string[];
        byte[] prevBlock = new byte[0];
        Blockchain block = new Blockchain(arg?[0] ?? "", prevBlock);
        for (int i = 0; i < arg?.Length; i++)
        {
            if (i >= 1)
            {
                block.AddBlock(arg?[i] ?? "");
            }
            var actualResult = System.Text.Encoding.ASCII.GetString(block.Blockchains[i].CurrentBlockHash);
            output.WriteLine($"block hash {i}: {actualResult}");
            Assert.Equal(expectedResult?.ToArray()[i], actualResult);
        }
    }
}