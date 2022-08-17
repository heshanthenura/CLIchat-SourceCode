import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.ArrayList;
import java.util.Scanner;

public class Server {
    public static void main(String[] args) throws IOException {
        //region Public Variables
        Scanner input = new Scanner(System.in);
        ServerSocket serverSocket = new ServerSocket();
        int PORT= 8080;
        //endregion

        //region Validate PORT
        while(true){
            try{
                System.out.print("Enter PORT Number:- ");
                PORT= Integer.parseInt(input.next());
                break;
            }catch (Exception e){
                PORT=8080;
                System.out.println("PORT not Valid");
            }
        }
        //endregion

        // Starting Server
        serverSocket = new ServerSocket(PORT);
        System.out.println("Server Started on PORT "+PORT);



        while (!serverSocket.isClosed()){
            Socket socket = serverSocket.accept();
            new Thread(new CLients(socket)).start();
        }
    }
}

class CLients implements Runnable{

    //region Public Vars
    public static ArrayList<CLients> clients = new ArrayList<>();
    private Socket socket;
    private String nickName;
    private PrintWriter pr;
    private BufferedReader br;
    //endregion

    public CLients(Socket socket) throws IOException {
        this.socket=socket;
        this.pr=new PrintWriter(socket.getOutputStream(),true);
        this.br=new BufferedReader(new InputStreamReader(socket.getInputStream()));
        try{
            this.nickName = br.readLine();
        }catch (Exception e) {
            System.out.println(e);
        }
        clients.add(this);
        System.out.println("SERVER: "+nickName+" connected to Chat");
        broadcastMessage("SERVER: "+this.nickName+" connected to Chat");

    }

    @Override
    //Read Messages
    public void run() {
        String clientMessage;
        while (socket.isConnected()){
            try {
                clientMessage=br.readLine();
                if (1==2){

                }else{
                    System.out.println("("+this.nickName+") "+clientMessage);
                    broadcastMessage("("+this.nickName+") "+clientMessage);
                }

            } catch (IOException e) {
                System.out.println("SERVER: "+this.nickName+" is Disconnected");
                broadcastMessage("SERVER: "+this.nickName+" is Disconnected");
                clients.remove(this);
                try {
                    this.socket.close();
                    this.br.close();
                    this.pr.close();
                } catch (IOException ex) {
                    ex.printStackTrace();
                }
                break;
            }
        }
    }

    //Broadcast Message
    public void broadcastMessage(String message){
        for (CLients c : clients) {
            if (!c.nickName.equals(nickName)){
                c.pr.println(message);
            }

        }
    }
}