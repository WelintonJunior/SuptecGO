export default function Form({ children, submitHandler }) {
  return <form onSubmit={submitHandler}>{children}</form>;
}
