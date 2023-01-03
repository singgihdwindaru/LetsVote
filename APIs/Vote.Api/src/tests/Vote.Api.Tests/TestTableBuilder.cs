
using Xunit.Abstractions;

namespace Vote.Api.Tests;

public class TestTableBuilder : IXunitSerializable
{
     public int Index { get; private set; }
    public string _testName { get; private set; }
    public TestTableBuilder()
    {
    }
    public TestTableBuilder(int indexTest, string testName)
    {
        Index = indexTest;
        _testName = testName;

        /* Note  : 
          Mock= mock; // cannot serialize/Deserialize Type Action delegate in current xunit v2.xx 
        */
    }

    public void Deserialize(IXunitSerializationInfo info)
    {
        try
        {
            Index = info.GetValue<int>("Index");
            _testName = info.GetValue<string>("TestName");
        }
        catch (Exception ex)
        {
            _ = ex.Message;
            throw;
        }
    }

    public void Serialize(IXunitSerializationInfo info)
    {
        try
        {
            info.AddValue("Index", Index);
            info.AddValue("TestName", _testName);
        }
        catch (Exception ex)
        {
            _ = ex.Message;
            throw;
        }
    }

    public override string ToString()
    {
        return _testName;
    }
}