namespace Vote.Api.Tests;

public class TestTable 
{
    public TestTable()
    {
        TestName = "Fill test name";
    }
    public string TestName { get; set; }
    public object? Args { get; set; }
    public bool WantError { get; set; }//parameter for testing exception
    public object? ExpectedResult { get; set; }
    public Action Mock { get; set; }

     public static TheoryData<TestTableBuilder> BuildTestTable(TestTable[] tc)
    {
        TheoryData<TestTableBuilder> td = new TheoryData<TestTableBuilder>();
        for (int i = 0; i < tc.Length; i++)
        {
            td.Add(new TestTableBuilder(i, tc[i].TestName));
        }
        return td;
    }
}
