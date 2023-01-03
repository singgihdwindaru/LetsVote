namespace Vote.Api;

public class time
{
    public static Func<long> GetUnixTimeWithDelegate = () =>
    {
         return new DateTimeOffset(DateTime.UtcNow).ToUnixTimeSeconds();
    };
}