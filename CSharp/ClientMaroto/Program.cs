using Grpc.Core;
using System;
using TestProto;

namespace ClientMaroto
{
    class Program
    {
        static void Main(string[] args)
        {
            var channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);

            var client = new Maroto.MarotoClient(channel);
            var user = "piroca, só q em c# agora!";

            var reply = client.TesteMaroto(new ObjetoEntrada { Nome = user });
            Console.WriteLine("Greeting: " + reply.Nome);



            channel.ShutdownAsync().Wait();
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}
