public class test {
    public static final String ANSI_GREEN_BACKGROUND = "\033[36m";
    public static final String ANSI_RESET = "\u001B[0m";
    public static void main(String[] args) {
        System.out.println(ANSI_GREEN_BACKGROUND + "This text has a green background but default text!" + ANSI_RESET);
    }
}
