namespace Vote.Api;

public class time
{
    public static Func<long> GetUnixTime = () =>
    {
         return new DateTimeOffset(DateTime.UtcNow).ToUnixTimeSeconds();
    };
}